import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { environment } from '../../environments/environment';
import { HttpHeaders, HttpParams } from '@angular/common/http';
import { map } from 'rxjs/operators';

@Injectable()
export class ShortenService {
    apiUrl: string;
    constructor(
        private http: HttpClient,
        // private jwtService: JwtService
    ) {
        this.apiUrl = environment.apiUrl
        if (this.apiUrl == '') {
            let loc = window.location;
            this.apiUrl = `http://${loc.hostname}:9090`;
        }

    }
    private formatErrors(error: any) {
            return throwError(error.error);
    }

    post(path: string, body: Object = {}): Observable<any> {
        return this.http.post(`${this.apiUrl}${path}`, body).pipe(catchError(this.formatErrors));
    }
    shorten(path: string, body: Object = {}) {
        return this.post(path, body).pipe(map(data => { return data }));
    }
}