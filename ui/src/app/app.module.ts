import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';

import { NgZorroAntdModule, NZ_LOCALE, enUS } from 'ng-zorro-antd';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { AuthModule } from './auth/auth.module';


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    NgZorroAntdModule.forRoot(),
    AuthModule,
    AppRoutingModule
  ],
  providers: [{provide: NZ_LOCALE, useValue: enUS}],
  bootstrap: [AppComponent]
})
export class AppModule { }
