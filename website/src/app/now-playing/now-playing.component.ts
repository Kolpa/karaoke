import { Component, OnInit } from '@angular/core';
import { SocketService } from '../socket.service';

@Component({
  selector: 'app-now-playing',
  templateUrl: './now-playing.component.html',
  styleUrls: ['./now-playing.component.css']
})
export class NowPlayingComponent implements OnInit {
  currentSong: string;

  constructor(private socketService: SocketService) {
    this.currentSong = null;
  }

  ngOnInit() {
    this.socketService.currentSong$.subscribe(song => {
      this.currentSong = song;
    });

    setInterval(() => {
      this.socketService.currentSong();
    }, 1000);
  }

  skipSong() {
    this.socketService.skipSong();
  }

}
