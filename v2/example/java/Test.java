package com.test;

import java.io.DataInputStream;
import java.io.File;
import java.io.FileInputStream;

import com.tabtoy.*;
import com.tabtoy.table.Config;

public class Test {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		System.out.println("aaaaa");
		
		
		File file = new File("C:\\Users\\weiwe\\eclipse-workspace\\testTabtoy\\res\\ttt.bytes");
		try {
			FileInputStream streamf = new FileInputStream(file);
			
			
			
			DataInputStream stream = new DataInputStream(streamf);
			
			
			DataReader reader = new DataReader(stream);
			if(!reader.ReadHeader(null)) {
				System.out.println("No Header");
			}
			
			Config instance = new Config();
			Config.getConfigDeserializeHandler().Deserialize(instance, reader);
			
			String a = instance.GetSampleByID(101l).Name;
			int size = instance.Sample.size();
			System.out.println(size);
		}catch(Exception e){
			System.out.println(e.getMessage());
		}
		
		
		
	}

}
