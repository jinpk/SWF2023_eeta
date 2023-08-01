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
import { fDateTime } from '../../../utils/formatTime';
// ----------------------------------------------------------------------

const StyledProductImg = styled('img')({
  top: 0,
  width: '100%',
  height: '100%',
  objectFit: 'cover',
  position: 'absolute',
});

// ----------------------------------------------------------------------

StoCard.propTypes = {
  sto: PropTypes.object,
};

export default function StoCard({ sto }) {
  const { id, billboard_id, organizer_address, organizer_url, organizer_image_url, name, start, end, user_share, organizer_share, goal, funded, sto_address} = sto;
  return (
    <Card>
      <Link to={`/dashboard/sto/${id}`}>
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
            {billboard_id}
          </Label>
        <StyledProductImg alt={name} src='/assets/images/covers/cover_1.jpg' />
      </Box>

      <Stack spacing={2} sx={{ p: 3 }}>
        <Link color="inherit" underline="hover">
          <Typography variant="subtitle2" noWrap>
            {name}
          </Typography>
        </Link>
        <Typography variant="subtitle1">
        모금액 : {funded.amount}
          </Typography>
          <Typography variant="subtitle2" noWrap>
          목표 금액 : {goal.amount}
          </Typography>
        <Stack direction="row" alignItems="center" justifyContent="space-between">
        
          <Typography variant="subtitle2">
          {new Date(start * 1000).toString()} ~ {new Date(end * 1000).toString()}
          </Typography>
        </Stack>
      </Stack>
      </Link>
    </Card>
  );
}
