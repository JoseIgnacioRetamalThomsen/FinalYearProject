import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { UserAuth } from 'src/models/userAuth';


const authUserURL = "http://35.197.242.214:9371/user";
const authLoginURL= "http://35.197.242.214:9371/login";

const httpOptions = {
  headers: new HttpHeaders({
   // 'token':'sdsdfas'
  //   'Content-Type': 'application/json',
  // 'Accept': '*/*',
  // 'Access-Control-Allow-Origin': '*',
  // 'Access-Control-Request-Method': 'OPTIONS'
    })
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  /** POST: add a new user to the database */
  createUser (user: UserAuth): Observable<UserAuth> {
    return this.http.post<UserAuth>(authUserURL, user, httpOptions)
   
  }

  /** POST: Login user. */
  loginUser(user: UserAuth): Observable<UserAuth> {
    return this.http.post<UserAuth>(authLoginURL, user, httpOptions)
  }
}
