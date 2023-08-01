import { Helmet } from 'react-helmet-async';
import { faker } from '@faker-js/faker';
import { useParams} from 'react-router-dom';
// @mui
import { useTheme } from '@mui/material/styles';
import { Grid, Container, Typography, Stack, Button } from '@mui/material';
// components
import Iconify from '../components/iconify';
// sections
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

// ----------------------------------------------------------------------

export default function DetailBillboardPage() {
  const { id } = useParams();
  const theme = useTheme();


  const handleClick = async () => {
    const accounts = await window.cosmostation.cosmos.request({
      method: "cos_account",
      params: { chainName: "eeta" },
    });       

    const aa = await fetch(`http://3.37.36.76:1317/cosmos/auth/v1beta1/accounts/${accounts.address}`);
    const account = await aa.json()
const seq = account.account.sequence
const an = account.account.account_number
    const response = await window.cosmostation.cosmos.request({
      method: "cos_signAmino",
      params: {
        chainName: "eeta",
        doc: {
          chain_id: "eeta",
          fee: { amount: [{ denom: "stake", amount: "5000" }], gas: "200000" },
          memo: "",
          msgs: [
            {
              type: "/eeta.billboard.MsgCreateBillboard",
              value: {
                board_type: "online",
                name: "Naver",
                description: "Naver",
                url: "https://naver.com",
                final_bid_price_per_minute: [{ denom: "krw", amount: "1000000" }],
              },
            },
          ],
          sequence: seq,
          account_number: an,
        },
      },
    });
  };
  return (
    <>
      <Helmet>
        <title> Dashboard | Minimal UI </title>
      </Helmet>

      <Container maxWidth="xl">
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={5}>
          <Typography variant="h4" gutterBottom>
            Billboard
          </Typography>
          <Stack direction="row" spacing={1}>
          <Button variant="contained" startIcon={<Iconify icon="eva:plus-fill" />} onClick={handleClick}>
            STO 모금하기
          </Button>
          <Button variant="contained" startIcon={<Iconify icon="eva:plus-fill" />} onClick={handleClick}>
            광고 입찰하기
          </Button>
          </Stack>
        </Stack>
        <Typography variant="h4" gutterBottom>
          DDP A1 구역 전면광고
        </Typography>
        <Grid container spacing={3}>
          <Grid item xs={12} md={6} lg={4}>
            <AppOrderTimeline
              title="펀딩 일정"
              list={[...Array(5)].map((_, index) => ({
                id: faker.datatype.uuid(),
                title: [
                  '1983, orders, $4220',
                  '12 Invoices have been paid',
                  'Order #37745 from September',
                  'New order placed #XF-2356',
                  'New order placed #XF-2346',
                ][index],
                type: `order${index + 1}`,
                time: faker.date.past(),
              }))}
            />
          </Grid>
          <Grid item xs={12} md={6} lg={8}>
            <AppWebsiteVisits
              title="광고 단가"
              subheader="(+43%) than last year"
              chartLabels={[
                '01/01/2003',
                '02/01/2003',
                '03/01/2003',
                '04/01/2003',
                '05/01/2003',
                '06/01/2003',
                '07/01/2003',
                '08/01/2003',
                '09/01/2003',
                '10/01/2003',
                '11/01/2003',
              ]}
              chartData={[
                {
                  name: 'Team A',
                  type: 'column',
                  fill: 'solid',
                  data: [23, 11, 22, 27, 13, 22, 37, 21, 44, 22, 30],
                },
                {
                  name: 'Team B',
                  type: 'area',
                  fill: 'gradient',
                  data: [44, 55, 41, 67, 22, 43, 21, 41, 56, 27, 43],
                },
                {
                  name: 'Team C',
                  type: 'line',
                  fill: 'solid',
                  data: [30, 25, 36, 30, 45, 35, 64, 52, 59, 36, 39],
                },
              ]}
            />
          </Grid>

          <Grid item xs={12} md={6} lg={8}>
            <AppConversionRates
              title="투자자 목록"
              subheader="총 112명"
              chartData={[
                { label: 'Italy', value: 400 },
                { label: 'Japan', value: 430 },
                { label: 'China', value: 448 },
                { label: 'Canada', value: 470 },
                { label: 'France', value: 540 },
                { label: 'Germany', value: 580 },
                { label: 'South Korea', value: 690 },
                { label: 'Netherlands', value: 1100 },
                { label: 'United States', value: 1200 },
                { label: 'United Kingdom', value: 1380 },
              ]}
            />
          </Grid>

          <Grid item xs={12} md={6} lg={4}>
            <AppTrafficBySite
              title="위험지수"
              list={[
                {
                  name: 'FaceBook',
                  value: 323234,
                  icon: <Iconify icon={'eva:facebook-fill'} color="#1877F2" width={32} />,
                },
                {
                  name: 'Google',
                  value: 341212,
                  icon: <Iconify icon={'eva:google-fill'} color="#DF3E30" width={32} />,
                },
                {
                  name: 'Linkedin',
                  value: 411213,
                  icon: <Iconify icon={'eva:linkedin-fill'} color="#006097" width={32} />,
                },
                {
                  name: 'Twitter',
                  value: 443232,
                  icon: <Iconify icon={'eva:twitter-fill'} color="#1C9CEA" width={32} />,
                },
              ]}
            />
          </Grid>
        </Grid>
      </Container>
    </>
  );
}
