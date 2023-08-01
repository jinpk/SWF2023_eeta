import { Helmet } from 'react-helmet-async';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
// @mui
import { Container, Stack, Typography, Button } from '@mui/material';
// mock
import PRODUCTS from '../_mock/sto';
import Iconify from '../components/iconify';
import StoList from '../sections/@dashboard/sto/StoList';

// ----------------------------------------------------------------------

export default function StoPage() {
  const [stos, setStos] = useState([]);

  useEffect(() => {
    async function fetchStos() {
      try {
        const response = await fetch('http://3.37.36.76:1317/eeta/sto/list_all_sto');
        const data = await response.json();
        setStos(data.stos);
      } catch (error) {
        console.error('Error fetching billboards:', error);
      }
    }

    fetchStos();
  }, []);
  return (
    <>
      <Helmet>
        <title> Dashboard: Products | Minimal UI </title>
      </Helmet>

      <Container>
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={5}>
          <Typography variant="h4" gutterBottom>
            모금중인 광고 STO
          </Typography>
        </Stack>

        <StoList stos={stos} />
      </Container>
    </>
  );
}
