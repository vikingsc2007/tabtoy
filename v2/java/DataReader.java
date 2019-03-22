package com.tabtoy;

import java.io.DataInputStream;

public class DataReader
{



    final int CombineFileVersion = 4;
    final int SIZE_INT = Integer.SIZE/8;
    final int SIZE_LONG = Long.SIZE/8;
    final int SIZE_FLOAT = Float.SIZE/8;
    final int SIZE_BOOL = 1;
    final int SIZE_CHAR = 2;

    LittleEndianInputStream _reader;
    long _boundPos;

    public boolean ConvertNewLine;
    
    public DataReader(DataInputStream stream ) throws Exception
    {
        this._reader = new LittleEndianInputStream(stream);
        this._boundPos = this._reader.available();
    }

    public DataReader(LittleEndianInputStream stream, long boundpos)
    {
        this._reader = stream;
        this._boundPos = boundpos;
    }

    public DataReader(DataReader reader, long boundpos)
    {
        this._reader = reader._reader;
        this._boundPos = boundpos;
        ConvertNewLine = reader.ConvertNewLine;
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
        return _reader.pos + size <= this._boundPos;
    }



    public boolean ReadHeader(String expectBuildID)  throws Exception
    {
    	String tag = this.ReadString();
        if (!tag.contentEquals("TT") )
        {
            return false;
        }

        int ver = ReadInt32();
        if (ver != CombineFileVersion)
        {
            return false;
        }

        String buildID = ReadString();
        if (expectBuildID != null && !expectBuildID.contentEquals(buildID))
        {
            return false;
        }
        
        // 文件校验码
        String fileCheckSum = ReadString();

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

        this.ConsumeData(len);

        byte[] chars = new byte[len];
//        for(int i = 0 ;i < len; i++){
//            chars[i] = _reader.readBytes();
//        }
        
        _reader.read(chars);
        _reader.pos+=len;
        String a= new String(chars,"UTF-8");
        
        return a;
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
        
        handler.Deserialize(element, new DataReader(this,_reader.pos+bound));
        return element;
    }

};
