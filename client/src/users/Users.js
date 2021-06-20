import React, { Component } from 'react';
import axios from 'axios';
import './Users.css';

const API_URL = '/api';

class Users extends Component{
    constructor(props){
        super(props);
        this.state = {
            isFetching: false,
            users: []
        };
    }
    componentDidMount(){
        this.fetchUsers();
        //this.timer = setInterval(() => this.fetchUsers(), 10000);
    }

    componentWillUnmount(){
        clearInterval(this.timer);
        this.timer = null;
    }

    fetchUsers = () => {
        this.setState({...this.state, isFetching: true});
        axios.get(`${API_URL}/users/`).then(response => {
            console.log(response);
            this.setState({users: response.data, isFetching: false})
        })
        .catch(e => {
            console.log(e);
            this.setState({...this.state, isFetching: false});
        });
    };

    render() {
        return(
            <div>
                <table class="styled-table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Email</th>
                        <th>Phone number</th>
                    </tr>
                </thead>
                <tbody>
                    {this.state.users.map((user) => {
                        return(
                            <tr>
                                <td>{user.name}</td>
                                <td>{user.email}</td>
                                <td>{user.phone_number}</td>
                            </tr>
                        )
                    })}
                </tbody>
                </table>
                <p>{this.state.isFetching ? 'Fetching users...' : ''}</p>
            </div>
        )
    }
}

export default Users;