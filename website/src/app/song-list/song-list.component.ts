import { Component, OnInit } from '@angular/core';
import { SocketService } from '../socket.service';

@Component({
  selector: 'app-song-list',
  templateUrl: './song-list.component.html',
  styleUrls: ['./song-list.component.css']
})
export class SongListComponent implements OnInit {
  songs: string[];

  constructor(private socketService: SocketService) {
    this.songs = [];
  }

  ngOnInit() {
    this.socketService.listSongs$.subscribe(msg => {
        this.songs = msg;
    });

    this.socketService.listSongs();
  }

  refresh() {
    this.socketService.listSongs();
  }

  onSongClick(song: string) {
    this.socketService.addSong(song);
  }
}
