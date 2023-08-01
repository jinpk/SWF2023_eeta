import {useState, useEffect} from 'react';
import { Helmet } from 'react-helmet-async';
import { Grid, Container, Typography } from '@mui/material';
import BILLBOARD from '../_mock/billboard';
import BillboardList from '../sections/@dashboard/billboard/BillboardList';

export default function BillboardPage() {
  const [billboards, setBillboards] = useState([]);

  useEffect(() => {
    async function fetchBillboards() {
      try {
        const response = await fetch('http://10.0.14.188:1317/eeta/billboards');
        const data = await response.json();
        setBillboards(data.billboard);
      } catch (error) {
        console.error('Error fetching billboards:', error);
      }
    }

    fetchBillboards();
  }, []);

  return (
    <>
      <Helmet>
        <title> Dashboard: Blog | Minimal UI </title>
      </Helmet>

      <Container>
          <Typography variant="h4" mb={5}>
            광고 목록
          </Typography>

        
        <BillboardList billboards={billboards} />
      </Container>
    </>
  );
}
