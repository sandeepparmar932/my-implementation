package com.sandeep.designPattern;


interface  Logger { void log(String msg);}

class ConsoleLogger implements Logger {
    @Override
    public void log(String msg) {
        System.out.println("ConsoleLogger - " + msg);
    }
}

class DebugLogger implements Logger {
    @Override
    public void log(String msg) {
        System.out.println("DebugLogger - " + msg);
    }
}

class InfoLogger implements Logger {
    @Override
    public void log(String msg) {
        System.out.println("InfoLogger - " + msg);
    }
}

class LoggerFactory {
    public static Logger getLogger(String type) throws Exception {
          switch (type) {
             case "Console" : return new ConsoleLogger();
             case "Debug" : return new DebugLogger();
             case "Info"  : return new InfoLogger();
             default: throw new Exception("Error ");

        }
    }
}

public class FactoryPattern {
    public static void main(String[] args) throws Exception {
        Logger logger = LoggerFactory.getLogger("Console");

    }
}
