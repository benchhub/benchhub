import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { NgZorroAntdModule } from 'ng-zorro-antd';
import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    NgZorroAntdModule.forRoot(),
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
