/* eslint-disable camelcase */
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
// @mui
import { Box, Card, Typography, Stack } from '@mui/material';
import { styled } from '@mui/material/styles';
// utils
import { fCurrency } from '../../../utils/formatNumber';
// components
import Label from '../../../components/label';
import { ColorPreview } from '../../../components/color-utils';

// ----------------------------------------------------------------------

const StyledProductImg = styled('img')({
  top: 0,
  width: '100%',
  height: '100%',
  objectFit: 'cover',
  position: 'absolute',
});

// ----------------------------------------------------------------------

BillboardCard.propTypes = {
  billboard: PropTypes.object,
};

export default function BillboardCard({ billboard }) {
  const { id, name, board_type, owner_address, url, final_bid_price_per_minute } = billboard;

  return (
    <Card>
      <Link to={`/dashboard/billboard/${id}`}>
      <Box sx={{ pt: '100%', position: 'relative' }}>
          <Label
            variant="filled"
            color={'info'}
            sx={{
              zIndex: 9,
              top: 16,
              right: 16,
              position: 'absolute',
              textTransform: 'uppercase',
            }}
          >
            {board_type}
          </Label>
        <StyledProductImg alt={name} src='/assets/images/covers/cover_4.jpg' />
      </Box>

      <Stack spacing={2} sx={{ p: 3 }}>
        <Link color="inherit" underline="hover">
          <Typography variant="subtitle2" noWrap>
            {name}
          </Typography>
        </Link>

        <Stack direction="row" alignItems="center" justifyContent="space-between">
          <Typography variant="subtitle1">
            최저 STO 진행 금액 : {final_bid_price_per_minute.amount} {final_bid_price_per_minute.denom}
          </Typography>
        </Stack>
      </Stack>
      </Link>
    </Card>
  );
}
