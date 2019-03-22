package com.tabtoy;

import java.io.DataInputStream;

class DataReader
{



    final int CombineFileVersion = 2;
    final int SIZE_INT = Integer.SIZE/8;
    final int SIZE_LONG = Long.SIZE/8;
    final int SIZE_FLOAT = Float.SIZE/8;
    final int SIZE_BOOL = 1;
    final int SIZE_CHAR = 2;

    DataInputStream _reader;



    public DataReader(DataInputStream stream ) throws Exception
    {
        this._reader = stream;
        this._reader.reset();
    }

    public DataReader(DataInputStream stream, long boundpos)
    {
        this._reader = stream;
    }

    public DataReader(DataReader reader)
    {
        this._reader = reader._reader;
    }

    private void ConsumeData(int size) throws Exception
    {
        if (!IsDataEnough( size ) )
        {
            throw new Exception("Out of struct bound");
        }
    }

    protected boolean IsDataEnough(int size) throws Exception
    {
        return size <= this._reader.available();
    }



    public boolean ReadHeader()  throws Exception
    {
        String tag = this.ReadString();
        if (!tag.contentEquals("TABTOY"))
        {
            return false;
        }

        int ver = this.ReadInt32();
        if (ver != this.CombineFileVersion)
        {
            return false;
        }

        return true;
    }

    public int ReadTag() throws Exception
    {
        if ( this.IsDataEnough(this.SIZE_INT))
        {
            return this.ReadInt32();
        }
        return -1;
    }

    public int ReadInt32()  throws Exception
    {
        this.ConsumeData(this.SIZE_INT);
        return _reader.readInt();
    }

    public long ReadInt64() throws Exception
    {
        this.ConsumeData(this.SIZE_LONG);
        return _reader.readLong();
    }

    public int ReadUInt32() throws Exception
    {
        this.ConsumeData(this.SIZE_LONG);
        return _reader.readInt()&0x0FFFFFFFF;
    }

    public long ReadUInt64() throws Exception
    {
        this.ConsumeData(this.SIZE_LONG);
        return _reader.readLong();
    }

    public float ReadFloat() throws Exception
    {
        this.ConsumeData(this.SIZE_FLOAT);
        return _reader.readFloat();
    }

    public boolean ReadBool() throws Exception
    {
        this.ConsumeData(this.SIZE_BOOL);
        return _reader.readBoolean();
    }

    public String ReadString() throws Exception
    {
        int len = this.ReadInt32();

        this.ConsumeData(len * this.SIZE_CHAR);

        char[] chars = new char[len];
        for(int i = 0 ;i < len; i++){
            chars[i] = _reader.readChar();
        }
        return new String(chars);
    }

    public <T> T ReadEnum(Class<T> clz) throws Exception{
        return clz.getEnumConstants()[this.ReadInt32()];
    }

    public <T> T ReadStruct(Class<T> clz,DeserializeHandler<T> handler) throws Exception
    {

        int bound = _reader.readInt();
        T element;
        try{
            element = clz.newInstance();
        }catch (Exception e){
            throw new Exception("instante struct error");
        }

        handler.Deserialize(element, new DataReader(this));
        return element;
    }

};
