## Waveform Audio File Format
wavio is a package for reading from and writing to wav audio files

     RIFF('WAVE'
          <fmt-ck>            // Format
          [<fact-ck>]         // Fact chunk
          [<cue-ck>]          // Cue points
          [<playlist-ck>]     // Playlist
          [<assoc-data-list>] // Associated data list
          <wave-data> )       // Wave data

     <wave-data> → data( <bSampleData:Byte> ... )

