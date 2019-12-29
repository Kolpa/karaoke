import { Injectable } from '@angular/core';
import { webSocket, WebSocketSubject } from "rxjs/webSocket";
import { Observable } from 'rxjs';

const buildWSURL = () => (window.location.protocol === "https:" ? "wss://" : "ws://") + window.location.host + "/ws";

@Injectable({
  providedIn: 'root'
})
export class SocketService {

  private socket: WebSocketSubject<any>;
  public listSongs$: Observable<any>;
  public currentSong$: Observable<any>;
  public getQueue$: Observable<any>;

  constructor() {
    //this.socket = webSocket(buildWSURL());
    this.socket = webSocket("ws://localhost:8080");

    this.listSongs$ = this.socket.multiplex(() => {}, () => {}, val => val.Command === "list_songs");
    this.currentSong$ = this.socket.multiplex(() => {}, () => {}, val => val.Command === "current_song");
    this.getQueue$ = this.socket.multiplex(() => {}, () => {}, val => val.Command === "get_queue");
  }

  public addSong(song: string) {
    this.socket.next({
      Command: "add_song",
      Arguments: song
    });
  }

  public currentSong() {
    this.socket.next({
      Command: "current_song"
    });
  }

  public getQueue() {
    this.socket.next({
      Command: "get_queue"
    });
  }

  public listSongs() {
    this.socket.next({
      Command: "list_songs"
    });
  }

  public skipSong() {
    this.socket.next({
      Command: "skip_song"
    });
  }
}
