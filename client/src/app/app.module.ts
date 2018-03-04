import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { Http } from '@angular/http'


import { AppComponent } from './app.component';
import { MovieService } from './app.service'
@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [
    Http,
    MovieService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
