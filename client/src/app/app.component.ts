import { Component } from '@angular/core';
import { MovieService } from './app.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public movies: Array<any>

  constructor(movieService: MovieService) {
    movieService.getMovies().subscribe(m => {
      this.movies.push(m)
    })
  }

  title = 'app';
}
