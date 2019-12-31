import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";

import { IndexPageComponent } from "./index-page/index-page.component";
import { SongListComponent } from "./song-list/song-list.component";
import { NowPlayingComponent } from "./now-playing/now-playing.component";
import { QueueComponent } from "./queue/queue.component";
import { NoopAnimationsModule } from "@angular/platform-browser/animations";
import { MatTableModule } from "@angular/material/table";
import { MatPaginatorModule } from "@angular/material/paginator";
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatInputModule } from "@angular/material/input";

@NgModule({
  declarations: [
    IndexPageComponent,
    SongListComponent,
    NowPlayingComponent,
    QueueComponent
  ],
  imports: [
    BrowserModule,
    NoopAnimationsModule,
    MatTableModule,
    MatPaginatorModule,
    MatFormFieldModule,
    MatInputModule
  ],
  providers: [],
  bootstrap: [IndexPageComponent]
})
export class AppModule {}
