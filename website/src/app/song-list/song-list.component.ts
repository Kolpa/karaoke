import { Component, OnInit, ViewChild } from "@angular/core";
import { SocketService, SongInfo } from "../socket.service";
import { MatTableDataSource } from "@angular/material/table";
import { MatPaginator } from "@angular/material/paginator";

@Component({
  selector: "app-song-list",
  templateUrl: "./song-list.component.html",
  styleUrls: ["./song-list.component.css"]
})
export class SongListComponent implements OnInit {
  displayedColumns = ["artist", "title"];
  songDataSource: MatTableDataSource<SongInfo[]>;

  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;

  constructor(private socketService: SocketService) {
    this.songDataSource = new MatTableDataSource();
  }

  ngOnInit() {
    this.songDataSource.paginator = this.paginator;

    this.socketService.listSongs$.subscribe(msg => {
      this.songDataSource.data = msg;
    });

    this.socketService.listSongs();
  }

  applyFilter(filterValue: string) {
    this.songDataSource.filter = filterValue;
  }

  rowSelected(song: SongInfo) {
    this.socketService.addSong(song.path);
  }
}
