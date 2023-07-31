import { Helmet } from 'react-helmet-async';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
// @mui
import { Container, Stack, Typography, Button } from '@mui/material';
// mock
import PRODUCTS from '../_mock/products';
import Iconify from '../components/iconify';
import StoList from '../sections/@dashboard/sto/StoList';

// ----------------------------------------------------------------------

export default function StoPage() {
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
            모금중인 빌보드 STO
          </Typography>
        </Stack>

        <StoList products={PRODUCTS} />
      </Container>
    </>
  );
}
