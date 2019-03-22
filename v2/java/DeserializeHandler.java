package com.tabtoy;

public interface DeserializeHandler<T>{
    public void Deserialize(T ins, DataReader reader) throws Exception;
}