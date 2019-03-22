package com.tabtoy;

import com.tabtoy.LogTarget;

class DebuggerTarget extends LogTarget
{
    @Override
    public void WriteLog(LogLevel level, String msg)
    {
        System.out.println( LevelToString(level) + msg );
    }
}