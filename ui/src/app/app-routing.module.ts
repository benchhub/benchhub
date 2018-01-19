import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// TODO: redirect based on local session, login or dashboard
const routes: Routes = [];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
