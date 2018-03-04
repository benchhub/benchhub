import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';

import { NgZorroAntdModule, NZ_LOCALE, enUS } from 'ng-zorro-antd';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { AuthModule } from './auth/auth.module';
import { DashModule } from './dash/dash.module';
import {HttpClientModule} from "@angular/common/http";


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    AuthModule,
    DashModule,
    AppRoutingModule,
    NgZorroAntdModule.forRoot()
  ],
  providers: [{provide: NZ_LOCALE, useValue: enUS}],
  bootstrap: [AppComponent]
})
export class AppModule { }
