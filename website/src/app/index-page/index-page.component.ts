import { Component, OnInit } from "@angular/core";
import { BreakpointObserver, BreakpointState } from "@angular/cdk/layout";

@Component({
  selector: "index-page",
  templateUrl: "./index-page.component.html",
  styleUrls: ["./index-page.component.css"]
})
export class IndexPageComponent implements OnInit {
  showQueue = false;

  constructor(public breakpointObserver: BreakpointObserver) {}

  ngOnInit() {
    const layoutChanges = this.breakpointObserver.observe([
      "(min-width: 500px)"
    ]);

    layoutChanges.subscribe((state: BreakpointState) => {
      this.showQueue = state.matches;
    });
  }
}
