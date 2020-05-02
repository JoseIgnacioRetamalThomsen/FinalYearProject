import { Component, OnInit } from '@angular/core';
import {SessionService} from '../session.service'
import { UserAuth } from 'src/models/userAuth';

@Component({
  selector: 'app-top-bar',
  templateUrl: './top-bar.component.html',
  styleUrls: ['./top-bar.component.css']
})
export class TopBarComponent implements OnInit {

  isLoggedIn = true;
  userName = "Jose Retamal";

  user: UserAuth = { name: "pepe", email: "my@gmail.com" ,hashPassword:""};

  searchText;
  constructor(private session :SessionService) {

    this.isLoggedIn = session.isLogin();
    if (this.isLoggedIn) {
      this.user.email = session.getName();
    }

   }

  ngOnInit(): void {


  }

  search() {

  //  this.router.navigate(['result/' + this.searchText]);
  }

 
  //logout
  logOut() {

    this.session.logOut();
    
    //navigate to home and reload page
    //this.router.navigate(['']);
    window.location.reload();
  }
}
