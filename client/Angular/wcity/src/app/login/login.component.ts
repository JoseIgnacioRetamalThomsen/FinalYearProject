import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { NgForm } from "@angular/forms";

import {AuthService} from '../auth.service'
import {SessionService} from '../session.service'
import { UserAuth } from 'src/models/userAuth';
import { iif } from 'rxjs';



@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {



   //show/hide password
   hide = true;

   //form controls/validator
   //Create and add validotes to password Form Control : required,email
   email = new FormControl('', [Validators.required, Validators.email]);
   //values for max/min password length
   minPasswordLength = 6;
   maxPasswordLength = 16;
   //Create and add validotes to password Form Control : required,minLength, maxLength
   password = new FormControl('', [Validators.required, Validators.minLength(this.minPasswordLength), Validators.maxLength(this.maxPasswordLength)]);
   //values for max min name lengt
   minNameLength = 1;
   maxNameLength = 32
   //Create and add validotes to password Form Control : required,minLength, maxLength
   name = new FormControl('', [Validators.required, Validators.minLength(this.minNameLength), Validators.maxLength(this.maxNameLength)]);
 
   //error mesages
   wrongDataMessage = false;
   emailInUseMessage = true;
   signinErrorMessage = "";
   signUpErrorMessage = "";
  

  getErrorMessage() {
    if (this.email.hasError('required')) {
      return 'You must enter a value';
    }
    return this.email.hasError('email') ? 'Not a valid email' : '';
  }
  constructor(private authService:AuthService, private session :SessionService) { }

  ngOnInit(): void {
  }

/*
  * Sign up
  */
 onSignUp(): void {

  var temp;

  if (this.email.valid && this.password.valid && this.name.valid) {

    const newUser: UserAuth = {email:this.email.value.trim(),hashPassword: this.password.value.trim(),name:this.name.value.trim() } as UserAuth;
    //create user, email to upper case
    this.authService.createUser(newUser).subscribe((data) => {
      console.log("not error");
      console.log(data);
      

    },(err)=>{
      console.log(" error");
      console.log(err);
      if(err.status = 500){
        this.signUpErrorMessage = "There is a account with that email.";
      }
    });

  }

}//onSignUp end

/*
* Sign in
*/
onSignIn(): void {

  var response;

  //check if email/password valid
  if (this.email.valid && this.password.valid) {

    const newUser: UserAuth = {email:this.email.value.trim(),hashPassword: this.password.value.trim(),name:this.name.value.trim() } as UserAuth;

    this.authService.loginUser(newUser).subscribe(data => {

      console.log("not error");
      console.log(data);
      response = data;

    }, (err) => {
      if (err.status == 401) {

        this.wrongDataMessage = true;
        this.signinErrorMessage = "You've entered an incorrect username/password.";

      } else if (err.status == 0) {

        this.wrongDataMessage = true;
        this.signinErrorMessage = "Can't connect to server, please try again later.";
      }

    }, () => {

      if (response.isUser) {
        console.log(response);
        //login successfull

        //log in in session (local storate)
        this.session.logIn(this.email.value, response.token);


        // //navigate to home page
        // this.router.navigate(['']);
        // //reload 
        //  window.location.reload();


      } else {//wrong data

        //show wrong data message
        this.wrongDataMessage = true;

      }//(response.success)

    });


  }

}

/*
*form control methods
*/
getErrorMessageEmail() {
  return this.email.hasError('required') ? 'You must enter a value' :

    this.email.hasError('email') ? 'Not a valid email' :
      '';
}

getErrorMessagePassword() {

  return this.password.hasError('required') ? 'You must enter a value' :

    this.password.hasError('minlength') ? 'Min ' + this.minPasswordLength + ' characters long' :

      this.password.hasError('maxlength') ? 'Max ' + this.maxPasswordLength + ' characters long' :

        '';

}

getErrorMessageName() {
  return this.name.hasError('required') ? 'You must enter a value' :

    this.name.hasError('minlength') ? 'Min ' + this.minNameLength + ' characters long' :

      this.name.hasError('maxlength') ? 'Max ' + this.maxNameLength + ' characters long' :

        '';
}

         






}


