// import { Helmet } from 'react-helmet-async';
// import { useState } from 'react';
// import { useNavigate } from 'react-router-dom';
// // @mui
// import { Container, Stack, Typography, Button } from '@mui/material';
// // components
// import { ProductSort, ProductList, ProductCartWidget, ProductFilterSidebar } from '../sections/@dashboard/sto';
// // mock
// import PRODUCTS from '../_mock/products';
// import Iconify from '../components/iconify';

// // ----------------------------------------------------------------------

// export default function BidPage() {
//   const navigate = useNavigate();
//   const [openFilter, setOpenFilter] = useState(false);

//   const handleOpenFilter = () => {
//     setOpenFilter(true);
//   };

//   const handleCloseFilter = () => {
//     setOpenFilter(false);
//   };

//   const handleClick = () => {
//     navigate('/new-npo', { replace: true });
//   };

//   return (
//     <>
//       <Helmet>
//         <title> Dashboard: Products | Minimal UI </title>
//       </Helmet>

//       <Container>
//         <Stack direction="row" alignItems="center" justifyContent="space-between" mb={5}>
//           <Typography variant="h4" gutterBottom>
//             NPO
//           </Typography>
//           <Button variant="contained" startIcon={<Iconify icon="eva:plus-fill" />} onClick={handleClick}>
//             New NPO
//           </Button>
//         </Stack>

//         <Stack direction="row" flexWrap="wrap-reverse" alignItems="center" justifyContent="flex-end" sx={{ mb: 5 }}>
//           <Stack direction="row" spacing={1} flexShrink={0} sx={{ my: 1 }}>
//             <ProductFilterSidebar
//               openFilter={openFilter}
//               onOpenFilter={handleOpenFilter}
//               onCloseFilter={handleCloseFilter}
//             />
//             <ProductSort />
//           </Stack>
//         </Stack>

//         <ProductList products={PRODUCTS} />
//         <ProductCartWidget />
//       </Container>
//     </>
//   );
// }
