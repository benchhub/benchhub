import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

// TODO: redirect based on local session to login or dashboard, now we are all redirected to dashboard
const routes: Routes = [
  {
    path: '',
    redirectTo: '/dash',
    pathMatch: 'full'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
