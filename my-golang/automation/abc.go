package main

/*
type StructInfo struct {
	Name   string
	Fields []FieldInfo
}

type FieldInfo struct {
	Name string
	Type string
	Tags string
}

type validateGormFields struct {
	name   string
	key    map[string]string
	values map[string]string
}

func (m *validateGormFields) Reset() {
	*m = validateGormFields{}
}

var ContainerList []string

func runPyangCommand(yangFile string, printReq bool) error {
	yangData, processedYangContainerList, err := validateYangFile(yangFile)
	if err != nil {
		fmt.Println("error while validating yang file", err.Error())
	}
	ContainerList = processedYangContainerList
	targetDir := os.Getenv("modelDir")
	fmt.Println("os", targetDir)

	if err := os.Chdir(targetDir); err != nil {
		return fmt.Errorf("%v", err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println("current dir", currentDir)
	for index, modelName := range yangData {
		if index == 0 {
			continue
		}
		fileName := ContainerList[index-1] + ".go"
		structName := ContainerList[index-1]

		currentStruct, err := ParseStruct("", fileName, structName)

		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Println("Matched", index, modelName.name, currentStruct.Name, ContainerList[index-1])
		if ContainerList[index-1] == currentStruct.Name {
			fmt.Println("Validating struct for : ", currentStruct.Name)

			for _, structFields := range currentStruct.Fields {
				checkGormValidation(modelName, structFields, currentStruct.Name)
			}
		}
	}

	return nil
}

func Run(caseType string, printReq bool) error {

	var targetDir string
	switch caseType {
	case "asset":
		targetDir = "../../../../../../ConfigState/etc/yang/asset"
	case "infer":
		targetDir = "../yang/infer"
	default:
		return fmt.Errorf("Error: Unsupported case %s", caseType)
	}

	if err := os.Chdir(targetDir); err != nil {
		return fmt.Errorf("Error changing directory to %s for case %s: %v", targetDir, caseType, err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting current working directory for case %s: %v", caseType, err)
	}

	if printReq {
		fmt.Printf("Current working directory for case %s: %s\n", caseType, currentDir)
	}

	yangFile := fmt.Sprintf("%s.yang", caseType)

	if err := runPyangCommand(yangFile, printReq); err != nil {
		return fmt.Errorf("Error running pyang command for case %s: %v", caseType, err)
	}

	return nil
}

func ParseStruct(folderPath, fileName, structName string) (*StructInfo, error) {
	filePath := filepath.Join(folderPath, fileName)

	fmt.Printf("Trying to read file: %s\n", filePath)

	src, err := io.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filePath, src, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}

	var foundStruct *StructInfo
	ast.Inspect(node, func(n ast.Node) bool {

		typeSpec, ok := n.(*ast.TypeSpec)
		if ok && typeSpec.Name.Name == structName {
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				fields := make([]FieldInfo, len(structType.Fields.List))
				for i, field := range structType.Fields.List {
					for _, name := range field.Names {
						// Get the struct tags
						var tags string
						if field.Tag != nil {
							tags = field.Tag.Value
						}

						fields[i] = FieldInfo{
							Name: name.Name,
							Type: fmt.Sprintf("%s", field.Type),
							Tags: tags,
						}
					}
				}
				foundStruct = &StructInfo{
					Name:   structName,
					Fields: fields,
				}
			}
		}
		return true
	})

	if foundStruct == nil {
		return nil, fmt.Errorf("struct %s not found in file %s", structName, filePath)
	}

	return foundStruct, nil
}

func BuildName(name string) string {
	ModelParts := strings.Split(name, "-")
	for i, part := range ModelParts {
		if len(part) > 0 {
			ModelParts[i] = strings.ToUpper(string(part[0])) + part[1:]
		}
	}
	return strings.Join(ModelParts, "")
}
func ExtractGormTags(input string) string {
	// Define a regex pattern to find the gorm tags
	re := regexp.MustCompile(`gorm:"([^"]+)"`)
	matches := re.FindStringSubmatch(input)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func checkGormValidation(model validateGormFields, structFields FieldInfo, structName string) error {

	if gormFieldValue, ok := model.values[structFields.Name]; ok {
		gormTags := ExtractGormTags(structFields.Tags)
		//fmt.Println("reflect.", reflect.TypeOf(gormTags))
		fmt.Println("in gorm validation go values", structFields.Name, gormTags, "yang values", gormFieldValue)

		if structFields.Type == "float64" {
			structFields.Type = "decimal64"
		}
		val := checkExists(gormTags, gormFieldValue)
		if val {
			log.Printf("Passed for Container Name: %s , Variable Name:= %s  \n", structName, structFields.Name)
		} else {
			log.Printf("Failed for Container Name: %s , Variable Name:=%s  \n", structName, structFields.Name)
		}
	}
	return nil
}

func validateYangFile(yangFile string) ([]validateGormFields, []string, error) {

	var moduleName string
	var ContainerList []string
	var ModelNames []validateGormFields
	var abc string
	args := []string{"-f", "tree", yangFile, "-p ~/xcp/PlatformServices/ConfigState/etc/yang/"}
	cmd := exec.Command("/home/user/.local/bin/pyang", args...)
	fmt.Println("cmd", cmd)

	cmdOutput, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating StdoutPipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("error starting command: %v", err)
	}

	validate := validateGormFields{}

	scanner := bufio.NewScanner(cmdOutput)

	mapOfIdKey := make(map[string]string)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "module") {
			parts := strings.Fields(line)
			moduleName = parts[1]
		}

		if validate.key == nil {
			validate.values = make(map[string]string)
		}

		if strings.HasPrefix(line, "+--rw") {
			ModelNames = append(ModelNames, validate)
			validate.Reset()

			parts := strings.Split(line, " ")
			if len(parts) < 2 {
				return nil, nil, fmt.Errorf("invalid input")
			}

			name := strings.TrimSuffix(parts[1], "*")
			abc = name

			idKey := "not null;uniqueIndex:idx"

			pattern := `\[\s*([^\]]+)\s*\]`
			re := regexp.MustCompile(pattern)
			matches := re.FindStringSubmatch(line)

			if len(matches) == 2 {
				ids := strings.Fields(matches[1])

				for _, id := range ids {
					idKey += "_" + BuildName(id)
					if validate.key == nil {
						validate.key = make(map[string]string)
					}
					validate.key[id] = id
				}
			}

			goModelName := BuildName(name)
			file := strings.ToUpper(string(moduleName[0])) + moduleName[1:]
			ContainerList = append(ContainerList, file+goModelName)

			mapOfIdKey[name] = idKey
			validate.name = name

		} else if strings.HasPrefix(line, "|") {
			parts := strings.Fields(line)
			if len(parts) <= 3 {
				continue
			}

			if validate.values == nil {
				validate.values = make(map[string]string)
			}

			checkNull := strings.HasPrefix(parts[3], "->")
			key := parts[2]
			if strings.HasSuffix(key, "?") {
				continue
			}

			constraintName := BuildName(strings.Trim(key, "?"))
			gormKey, _ := mapOfIdKey[abc]

			_, ok := validate.key[key]
			if checkNull && !strings.HasSuffix(key, "?") && !ok {
				validate.values[constraintName] = gormKey + ";"
			} else if ok {
				validate.values[constraintName] = gormKey + ";"
			} else if !strings.HasSuffix(key, "?") {
				validate.values[constraintName] = "not null;"
			}

			if parts[3] == "ietf-yang:uuid" && !ok {
				validate.values[constraintName] = "primary_key;"
			} else if checkNull && !ok {
				validate.values[constraintName] = "not null;"
			} else {
				validate.values[constraintName] = gormKey + ";"
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading command output: %v", err)
	}
	fmt.Println("test", ModelNames[1])

	return ModelNames, ContainerList, nil
}

func main() {
	args := os.Args
	printReq := len(args) == 3 && args[2] == "-print"
	caseType := args[1]
	os.Setenv("modelDir", "../../../src/configstate/domain/models")
	fmt.Println("vvv", os.Getenv("modelDir"))
	if err := Run(caseType, printReq); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func checkExists(input1, input2 string) bool {

	prefix := "not null;uniqueIndex:idx_"

	cleanedString1 := strings.Replace(input1, prefix, "", 1)
	cleanedString2 := strings.Replace(input2, prefix, "", 1)
	trimmedString1 := strings.TrimSuffix(cleanedString1, ";")
	trimmedString2 := strings.TrimSuffix(cleanedString2, ";")
	parts1 := strings.Split(trimmedString1, "_")
	parts2 := strings.Split(trimmedString2, "_")

	exists := make(map[string]bool)

	for _, part := range parts1 {
		exists[part] = true
	}

	termsToCheck := make(map[string]bool)
	for _, part := range parts2 {
		termsToCheck[part] = true
	}

	for term := range termsToCheck {
		if !exists[term] {
			return false
		}
	}

	return true
}*/
