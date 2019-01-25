import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';


const CustomTableCell = withStyles(theme => ({
  head: {
    backgroundColor: theme.palette.common.black,
    color: theme.palette.common.white,
  },
  body: {
    fontSize: 14,
  },
}))(TableCell);

const styles = theme => ({
  root: {
    width: '100%',
    marginTop: theme.spacing.unit * 3,
    overflowX: 'auto',
  },
  table: {
    minWidth: 700,
  },
  row: {
    '&:nth-of-type(odd)': {
      backgroundColor: theme.palette.background.default,
    },
  },
});



function personDataObjectGenerator(id, Fullname, Age, Salary, Mobile, Email, Address) {
	console.log("​personData -> Firstname", Fullname)
  return {id, Fullname, Age, Salary, Mobile, Email, Address};
}


function CustomizedTable(props) {
  const { classes } = props;
  console.log("​CustomizedTable -> props", props)
  let personListData = null
  if (props.PersonObject) {
    personListData = props.PersonObject.map((personData) => {
      const ifContact = personData.Contact ? [personData.contact.mobile,
        personData.address] : [0, 0, 0]
        
      const params = [personData.ID,personData.full_name, personData.age, personData.slary,
        ...ifContact
      ]
      console.log("​CustomizedTable -> params", params)
     
      return personDataObjectGenerator(...params)
    })
  }
  
  console.log("​personListData -> personListData", personListData)

  return (
    <Paper className={classes.root}>

      <Table className={classes.table}>
        <TableHead>
          <TableRow>
            <CustomTableCell>Person ID</CustomTableCell>
            <CustomTableCell>Fullname</CustomTableCell>
            <CustomTableCell>Age</CustomTableCell>
            <CustomTableCell>Salary</CustomTableCell>
            <CustomTableCell>Mobile No</CustomTableCell>
            <CustomTableCell>Email Id</CustomTableCell>
            <CustomTableCell>Address</CustomTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {personListData ? personListData.map(row => {
            return (
              <TableRow className={classes.row} key={row.id}>
                <CustomTableCell component="th" scope="row">
                  {row.id}
                </CustomTableCell>
                <CustomTableCell>
                  {row.Firstname}
                </CustomTableCell>
                <CustomTableCell>{row.Age}</CustomTableCell>
                <CustomTableCell>{row.Gender}</CustomTableCell>
                <CustomTableCell>{row.Mobile}</CustomTableCell>
                <CustomTableCell>{row.Email}</CustomTableCell>
                <CustomTableCell>{row.Address}</CustomTableCell>
              </TableRow>
            );
          }):''}
        </TableBody>
      </Table>

    </Paper>
  );
}

CustomizedTable.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(CustomizedTable);
