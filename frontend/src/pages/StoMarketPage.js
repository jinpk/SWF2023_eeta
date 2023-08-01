import { Helmet } from 'react-helmet-async';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
// @mui
import { Container, Stack, Typography, Button } from '@mui/material';
// mock
import PRODUCTS from '../_mock/sto';
import Iconify from '../components/iconify';
import StoList from '../sections/@dashboard/sto/StoList';

// ----------------------------------------------------------------------

export default function StoMarketPage() {
  const navigate = useNavigate();
  const [openFilter, setOpenFilter] = useState(false);

  const handleOpenFilter = () => {
    setOpenFilter(true);
  };

  const handleCloseFilter = () => {
    setOpenFilter(false);
  };

  const handleClick = () => {
    navigate('/new-npo', { replace: true });
  };

  return (
    <>
      <Helmet>
        <title> Dashboard: Products | Minimal UI </title>
      </Helmet>

      <Container>
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={5}>
          <Typography variant="h4" gutterBottom>
            STO 마켓플레이스
          </Typography>
        </Stack>

        <StoList products={PRODUCTS} />
      </Container>
    </>
  );
}
