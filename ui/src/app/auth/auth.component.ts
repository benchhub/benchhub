import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-auth',
  template: `
    <div>I should be the logo on the left</div>
    <div>
      I should be the dynamic content on the right
      <router-outlet></router-outlet>
    </div>
  `,
  styles: []
})
export class AuthComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
