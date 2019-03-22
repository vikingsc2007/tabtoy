package com.tabtoy;

public class LogTarget
{
    public void WriteLog(LogLevel level, String msg)
    {

    }

    public static String LevelToString( LogLevel level )
    {
        switch (level)
        {
            case Debug:
                return "tabtoy [Debug] ";                    
            case Info:
                return "tabtoy [Info] ";                    
            case Warnning:
                return "tabtoy [Warn] ";
            case Error:
                return "tabtoy [Error] ";
        }
        return "tabtoy [Unknown] ";
    }
}