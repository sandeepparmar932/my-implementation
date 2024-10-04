package com.sandeep.thread;
import java.io.Serializable;

public class Singleton implements Serializable, Cloneable {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1L;
	private static Singleton obj;

	private Singleton() throws Exception {
		if (obj != null) {
			throw new Exception("trying wrong");
		}
	}

	public static Singleton getInstance() throws Exception {
		if (obj == null) {
			synchronized (Singleton.class) {
				if (obj == null) {
					obj = new Singleton();
				}
			}

		}
		return obj;
	}

	public Singleton readResolver() {
		return this.obj;
	}

	@Override
	protected Singleton clone() throws CloneNotSupportedException {
		return this.obj;
	}

}
