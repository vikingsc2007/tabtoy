package com.mixmarvel.tabtoy;

class Logger
{
    ArrayList<LogTarget> _targets = new ArrayList<LogTarget>();

    public void AddTarget(LogTarget tgt)
    {
        _targets.add(tgt);
    }

    public void ClearTargets( )
    {
        _targets.clear();
    }

    private void WriteLine(LogLevel level, String msg)
    {
        for (int i = 0 ;i< _targets.size() ;i++){
            _targets.get(i).WriteLog(level, msg);
        }
    }

    public void DebugLine(String fmt, Object... args)
    {
        String text = String.format(fmt, args);

        this.WriteLine(LogLevel.Debug, text);
    }

    public void InfoLine(String fmt, Object... args)
    {
        String text = String.format(fmt, args);

        this.WriteLine(LogLevel.Info, text);
    }

    public void WarningLine(String fmt, Object... args)
    {
        String text = String.format(fmt, args);

        this.WriteLine(LogLevel.Warnning, text);
    }

    public void ErrorLine(String fmt, Object... args)
    {
        String text = String.format(fmt, args);

        this.WriteLine(LogLevel.Error, text);
    }
}
