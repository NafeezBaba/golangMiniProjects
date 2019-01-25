import React from 'react';
// import { Field, reduxForm } from 'redux-form'
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
// import Radio from '@material-ui/core/Radio';
import RedioButton from './redio';
// import RadioGroup from '@material-ui/core/RadioGroup';
// import FormHelperText from '@material-ui/core/FormHelperText';
import FormControlLabel from '@material-ui/core/FormControlLabel';
// import FormControl from '@material-ui/core/FormControl';
// import FormLabel from '@material-ui/core/FormLabel';
import axios from 'axios'
var API_URL = 'http://localhost:12345'

export default class FormDialog extends React.Component {
  state = {
    open: false,
  };

  handleClickOpen = () => {
    this.setState({ open: true });
  };

  handleClose = () => {
    this.setState({ open: false });
  };

  PostPerson(e) {
	console.log("​FormDialog -> PostPerson -> e", e)
    const body = {
      Firstname: '',
      Lastname: '',
      Gender: '',
      Age: 24,
      Address: {
        City: '',
        State: '',
        Pin: '',
      },
      Contact: {
        Mobile: '',
        Email: ''
      }
    }

    axios.post(API_URL + '/person', body, {
      headers: {
        'content-type': 'application/json',
      },
    })
      .then(res => {
        console.log("​Form -> PostPerson -> data", body)
        console.log('TCL: response', res);
        return res;
      }).catch(res => res = {
        code: 500,
        message: "Your submission could not be completed. Please Try Again!",
        data: ""
      });
    this.setState({ massege: 'Person Posted' })
  }

  render() {
    return (
      <div>
        <Button style={{width:'full'}} variant="contained" color="primary"  onClick={this.handleClickOpen}>Post Person</Button>
        <Dialog
          open={this.state.open}
          onClose={this.handleClose}
          aria-labelledby="form-dialog-title"
        >  <form>
          <DialogTitle id="form-dialog-title">Subscribe</DialogTitle>
        
          <DialogContent>   
            <DialogContentText>
              To Post the person into Database.
            </DialogContentText>
          
            <TextField autoFocus margin="dense" id="FirstName" label="First Name"type="text"/>
            <TextField  margin="dense" id="LastName" label="Last Name"type="text"/>
            <TextField  margin="dense" id="Gender" label="Gender"type="text"/>
                      <RedioButton/>
            
            <TextField  margin="dense" id="Age" label="Age"type="number"/>
            <TextField  margin="dense" id="Mobile" label="Mobile Number"type="text"/>
            <TextField  margin="dense" id="Email" label="Email ID"type="text"/>
            
            <TextField  margin="dense" id="City" label="City"type="text"/>
            <TextField  margin="dense" id="State" label="State"type="text"/>
            <TextField  margin="dense" id="Pin" label="Pin"type="text"/>
          
            
            
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose} color="primary">
              Cancel
            </Button>
            <Button type="submit" onClick={this.PostPerson.bind(this)} color="primary">
              Subscribe
            </Button>
           
          </DialogActions>
          </form>
        </Dialog>
      </div>
    );
  }
}
