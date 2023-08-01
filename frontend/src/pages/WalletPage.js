/* eslint-disable camelcase */
import { Helmet } from 'react-helmet-async';
import { filter } from 'lodash';
import { sentenceCase } from 'change-case';
import { useState, useEffect } from 'react';
import { useTheme } from '@mui/material/styles';
import { faker } from '@faker-js/faker';
// @mui
import {
  Card,
  Table,
  Stack,
  Paper,
  Avatar,
  Button,
  Popover,
  Checkbox,
  TableRow,
  MenuItem,
  TableBody,
  Grid,
  TableCell,
  Container,
  Typography,
  IconButton,
  TableContainer,
  TablePagination,
} from '@mui/material';
// sections
import { useNavigate } from 'react-router-dom';
import {
  AppTasks,
  AppNewsUpdate,
  AppOrderTimeline,
  AppCurrentVisits,
  AppWebsiteVisits,
  AppTrafficBySite,
  AppWidgetSummary,
  AppCurrentSubject,
  AppConversionRates,
} from '../sections/@dashboard/app';
// components
import Label from '../components/label';
import Iconify from '../components/iconify';
import Scrollbar from '../components/scrollbar';
// sections
import { UserListHead, UserListToolbar } from '../sections/@dashboard/user';
// mock
import USERLIST from '../_mock/user';

// ----------------------------------------------------------------------

const TABLE_HEAD = [
  { id: 'name', label: 'Name', alignRight: false },
  { id: 'company', label: 'Company', alignRight: false },
  { id: 'role', label: 'Role', alignRight: false },
  { id: 'isVerified', label: 'Verified', alignRight: false },
  { id: 'status', label: 'Status', alignRight: false },
  { id: '' },
];

// ----------------------------------------------------------------------

function descendingComparator(a, b, orderBy) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

function getComparator(order, orderBy) {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function applySortFilter(array, comparator, query) {
  const stabilizedThis = array.map((el, index) => [el, index]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  if (query) {
    return filter(array, (_user) => _user.name.toLowerCase().indexOf(query.toLowerCase()) !== -1);
  }
  return stabilizedThis.map((el) => el[0]);
}

export default function WalletPage() {
  const navigate = useNavigate();

  const [open, setOpen] = useState(null);

  const [page, setPage] = useState(0);

  const [order, setOrder] = useState('asc');

  const [selected, setSelected] = useState([]);

  const [orderBy, setOrderBy] = useState('name');

  const [filterName, setFilterName] = useState('');

  const [rowsPerPage, setRowsPerPage] = useState(5);

  const [account, setAccount] = useState();

  const [balances, setBalances] = useState();

  const handleOpenMenu = (event) => {
    setOpen(event.currentTarget);
  };

  const handleCloseMenu = () => {
    setOpen(null);
  };

  const handleRequestSort = (event, property) => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };

  const handleSelectAllClick = (event) => {
    if (event.target.checked) {
      const newSelecteds = USERLIST.map((n) => n.name);
      setSelected(newSelecteds);
      return;
    }
    setSelected([]);
  };

  const handleClick = (event, name) => {
    const selectedIndex = selected.indexOf(name);
    let newSelected = [];
    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selected, name);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selected.slice(1));
    } else if (selectedIndex === selected.length - 1) {
      newSelected = newSelected.concat(selected.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(selected.slice(0, selectedIndex), selected.slice(selectedIndex + 1));
    }
    setSelected(newSelected);
  };

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setPage(0);
    setRowsPerPage(parseInt(event.target.value, 10));
  };

  const handleFilterByName = (event) => {
    setPage(0);
    setFilterName(event.target.value);
  };

  const handleDeposit = () => {
    navigate('/deposit', { replace: true });
  };

  const theme = useTheme();
  const emptyRows = page > 0 ? Math.max(0, (1 + page) * rowsPerPage - USERLIST.length) : 0;

  const filteredUsers = applySortFilter(USERLIST, getComparator(order, orderBy), filterName);


  useEffect(() => {
    async function fetchAccount() {
      try {
        const account = await window.cosmostation.cosmos.request({
          method: "cos_account",
          params: { chainName: "eeta" },
        }); 
      setAccount(account);
      } catch (error) {
        console.error('Error fetching billboards:', error);
      }
    }

    async function fetchStos() {
      try {
        const response = await fetch('http://3.37.36.76:1317/eeta/sto/list_all_sto');
        const data = await response.json();
        setStos(data.stos);
        console.log(data);
      } catch (error) {
        console.error('Error fetching billboards:', error);
      }
    }

    async function fetchBalances() {
      const accounts = await window.cosmostation.cosmos.request({
        method: "cos_account",
        params: { chainName: "eeta" },
      });       
      const b = await fetch(`http://3.37.36.76:1317/cosmos/bank/v1beta1/balances/${accounts.address}`);
      const data = await b.json();
      setBalances(data.balances);
      console.log(data.balances);
    }
    async function fetchBillboards() {
      try {
        const response = await fetch('http://3.37.36.76:1317/eeta/billboards');
        const data = await response.json();
        setBillboards(data.billboard);
      } catch (error) {
        console.error('Error fetching billboards:', error);
      }
    }
    fetchBalances();
    fetchBillboards();
    fetchStos();
    fetchAccount();
  }, []);
  
  const [billboards, setBillboards] = useState([]);
  const [stos, setStos] = useState([]);
   
  return (
    <>
      <Helmet>
        <title> User | Minimal UI </title>
      </Helmet>

      <Container>
      <Typography variant="h4" gutterBottom>
            Hello 🙈
          </Typography>
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={5}>
        {balances && <Typography variant="h1" component="h2">
            {balances.filter(i => i.denom === 'krw')[0].amount * 0.000001} KRW
          </Typography>          }
          <Button variant="contained" startIcon={<Iconify icon="eva:plus-fill" />} onClick={handleDeposit}>
            Deposit
          </Button>
        </Stack>


        <Grid container spacing={3}>

{balances && 
<Grid item xs={12} md={6} lg={6}>
  <AppNewsUpdate
    title="참여중인 펀딩"
    list={balances.map((item, index) => ({
      id: index,
      title:  item.amount,
      description: item.denom,
      image: `/assets/images/covers/cover_2.jpg`,
      postedAt: faker.date.recent(),
    }))}
  />
</Grid>
}
{billboards &&
<Grid item xs={12} md={6} lg={6}>
  <AppNewsUpdate
    title="나의 광고"
    list={billboards.map((item, index) => ({
      id: item.id,
      title: item.name,
      description: item.final_bid_price_per_minute.amount + item.board_type,
      image: `/assets/images/covers/cover_3.jpg`,
      postedAt: faker.date.recent(),
    }))}
  />
</Grid>
}

{stos && 
<Grid item xs={12} md={6} lg={6}>
  <AppNewsUpdate
    title="내가 모금중인 STO"
    list={stos.map((item, index) => ({
      id: item.id,
      title:  item.name,
      description: `목표:${  item.goal.amount  }, 모금:${  item.funded.amount}`,
      image: `/assets/images/covers/cover_4.jpg`,
      postedAt: faker.date.recent(),
    }))}
  />
</Grid>
}
{stos && 
<Grid item xs={12} md={6} lg={6}>
  <AppNewsUpdate
    title="내가 판매중인 STO"
    list={stos.map((item, index) => ({
      id: item.id,
      title:  item.name,
      description: `목표:${  item.goal.amount  }, 모금:${  item.funded.amount}`,
      image: `/assets/images/covers/cover_6.jpg`,
      postedAt: faker.date.recent(),
    }))}
  />
</Grid>
}
                    </Grid>
      </Container>
    </>
  );
}