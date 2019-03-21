package com.mixmarvel.tabtoy;

public interface DeserializeHandler<T>{
    public void callBackMethod(T ins, DataReader reader);
}
