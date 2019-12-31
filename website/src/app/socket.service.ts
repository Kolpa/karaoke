import { Injectable } from "@angular/core";
import { webSocket, WebSocketSubject } from "rxjs/webSocket";
import { Observable } from "rxjs";
import { map, filter } from "rxjs/operators";

const buildWSURL = () =>
  (window.location.protocol === "https:" ? "wss://" : "ws://") +
  window.location.host +
  "/ws";

@Injectable({
  providedIn: "root"
})
export class SocketService {
  private socket: WebSocketSubject<any>;
  public listSongs$: Observable<any>;
  public currentSong$: Observable<any>;
  public getQueue$: Observable<any>;

  constructor() {
    //this.socket = webSocket(buildWSURL());
    this.socket = webSocket("ws://localhost:8080");

    this.listSongs$ = this.socket.multiplex(
      () => ({
        Command: "list_songs"
      }),
      () => ({
        Command: "list_songs"
      }),
      val => val.Command === "list_songs"
    ).pipe(
      filter(resp => resp.Status === "ok"),
      map(resp => resp.Songs)
    );
    this.currentSong$ = this.socket.multiplex(
      () => ({
        Command: "current_song"
      }),
      () => ({
        Command: "current_song"
      }),
      val => val.Command === "current_song"
    ).pipe(
      filter(resp => resp.Status === "ok"),
      map(resp => resp.Song)
    );
    this.getQueue$ = this.socket.multiplex(
      () => ({
        Command: "get_queue"
      }),
      () => ({
        Command: "get_queue"
      }),
      val => val.Command === "get_queue"
    ).pipe(
      filter(resp => resp.Status === "ok"),
      map(resp => resp.Songs)
    );
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
