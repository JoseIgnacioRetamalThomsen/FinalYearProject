import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SessionService {

  constructor() {
       //create login if do not exist
       if (localStorage.getItem('isLogin') == null) {
        localStorage.setItem('isLogin', "false");
      }
   }

    //login, save all data in local storage 
  logIn( email: string, token: string) {

    localStorage.setItem('isLogin', "true");
        localStorage.setItem('email', email);
     localStorage.setItem('token', token);

  }

    //remove all data from local stora and set login to false
    logOut() {

      localStorage.setItem('isLogin', "false");
     
      localStorage.removeItem('email');
      
      localStorage.removeItem('token');
  
    }

      /*
  * Getters for all data
  */
  isLogin(): boolean {
    if (localStorage.getItem('isLogin') == "false")
      return false;
    return true;
  }
  getName(): string {
    return localStorage.getItem('name');
  }
  getEmail(): string {
    return localStorage.getItem('email').toUpperCase();
  }

  getToken(): string {
    return localStorage.getItem('token');
  }
  
}
