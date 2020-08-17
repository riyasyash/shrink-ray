import { Component, OnInit } from '@angular/core';
import { ShortenService } from './services';
import { environment } from '../environments/environment';
import { MatSnackBar } from '@angular/material/snack-bar';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  shrinked: boolean = false;
  url: string = null;
  key: string = null;
  constructor(
    private shortenSVC: ShortenService,
    private _snackBar: MatSnackBar
  ) { }

  ngOnInit() {
  }

  shrink() {
    if(!this.validUrl()){
      this._snackBar.open("Please enter a valid url", '', {
        duration: 5000,
      });
      return
    }
    this.shortenSVC.shorten('/api/shorten', { url: this.url }).subscribe(
      res => {
        this.key = `${environment.hostUrl}/${res}`;
        this.shrinked = true;
      },
      err => {
        console.log(err)
        // const errorMsg = err.errors[0] ? err.errors[0] : 'Error while Logging.';
      }
    );
  }
  validUrl(){
    if (!this.url){
      return false
    }
    if(this.url.startsWith('http://')||this.url.startsWith('https://')){
      return true;
    }
    return false;
  }
}
