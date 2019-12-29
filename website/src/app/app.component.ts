import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <!--The content below is only a placeholder and can be replaced.-->
    <div style="text-align:center" class="content">
      <h1>
        Welcome to {{title}}!
        <app-now-playing></app-now-playing>
      </h1>
    </div>
    <app-song-list></app-song-list>
  `,
  styles: []
})
export class AppComponent {
  title = 'website';
}
