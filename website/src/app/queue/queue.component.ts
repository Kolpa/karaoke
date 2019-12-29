import { Component, OnInit } from '@angular/core';
import { SocketService } from '../socket.service';

@Component({
  selector: 'app-queue',
  templateUrl: './queue.component.html',
  styleUrls: ['./queue.component.css']
})
export class QueueComponent implements OnInit {
  songs: string[];

  constructor(private socketService: SocketService) {
    this.songs = [];
  }

  ngOnInit() {
    this.socketService.getQueue$.subscribe(msg => {
      if (msg.Status === "ok") {
        this.songs = msg.Songs;
      }
    });

    setInterval(() => {
      this.socketService.getQueue();
    }, 1000);
  }
}
