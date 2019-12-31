import { Component, OnInit } from "@angular/core";
import { SocketService, SongInfo } from "../socket.service";

@Component({
  selector: "app-now-playing",
  templateUrl: "./now-playing.component.html",
  styleUrls: ["./now-playing.component.css"]
})
export class NowPlayingComponent implements OnInit {
  currentSong: SongInfo;

  constructor(private socketService: SocketService) {}

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
