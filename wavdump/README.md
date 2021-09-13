## wavdump
Display information from wav file

    Usage: wavdump [options] wavfile
    options:
      -F    print samples as IEEE floats
      -N string
            range of samples to print/plot
            examples:
              -N 100     first 100 samples
              -N 50:100  50th thru 100th sample
              -N 100:    from 100th sample to the end of the file
      -P    pretty print samples
      -f    plot fft(x) (range must be 2^N)
      -log float
            plot log(rms(x))
            examples:
              -log=-40   floor >= -40 dB
      -p    plot x
      -r    plot rms(x)
    environment variables:
      WAVDUMP=term=iterm xres=800 yres=200    terminal graphics (iTerm2/mintty)
      WAVDUMP=nogfx    disable graphics (Kitty terminal)

```
$ wavdump tone.wav
tone.wav: fmt 1 (PCM) 16-bit 1 ch 44100 Hz 88200 Bps 2 align 88200 bytes [0:800]
```
![wavdump plot](wavdump.png)
