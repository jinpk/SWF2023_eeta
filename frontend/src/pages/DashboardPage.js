import { Helmet } from 'react-helmet-async';
import { faker } from '@faker-js/faker';
// @mui
import { useTheme } from '@mui/material/styles';
import { Grid, Container, Typography } from '@mui/material';
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

export default function DashboardAppPage() {
  const theme = useTheme();

  return (
    <>
      <Helmet>
        <title> Dashboard | Minimal UI </title>
      </Helmet>

      <Container maxWidth="xl">
        <Typography variant="h4" sx={{ mb: 5 }}>
          Ads Adoption
        </Typography>
        <Typography variant="h5" sx={{ mb: 5 }}>
          우리 실생활에서 가장 많이 활용되는 광고에서의 Mass Adaption을 달성하기 위해 Ads Adoption 프로젝트를 만들었습니다.
        </Typography>

        <Grid container spacing={3}>
          <Grid item xs={12} sm={6} md={3}>
            <AppWidgetSummary title="등록된 광고수" total={933} icon={'ant-design:dashboard-filled'} />
          </Grid>

          <Grid item xs={12} sm={6} md={3}>
            <AppWidgetSummary title="모금중인 STO" total={20} color="info" icon={'ant-design:fund-filled'} />
          </Grid>

          <Grid item xs={12} sm={6} md={3}>
            <AppWidgetSummary title="총 모금액" total={100000000} color="warning" icon={'ant-design:bank-filled'} />
          </Grid>

          <Grid item xs={12} sm={6} md={3}>
            <AppWidgetSummary title="평균 수익율" total={42} color="error" icon={'ant-design:rocket-filled'} />
          </Grid>

          <Grid item xs={12} md={6} lg={8}>
            <AppConversionRates
              title="진행중인 STO"
              chartData={[
                { label: '서울시 구의역 지하철 광고', value: 80 },
                { label: '서울시청 광고', value: 87 },
                { label: '서장훈 전광판', value: 67 },
                { label: 'DDP a1구역', value: 95 },
                { label: '유투브 - 뻐꾸기 골프', value: 39 },
              ]}
            />
          </Grid>

          <Grid item xs={12} md={6} lg={4}>
            <AppCurrentVisits
              title="광고 현황"
              chartData={[
                { label: '유투브', value: 4344 },
                { label: '온라인', value: 5435 },
                { label: '옥외광고', value: 1443 },
                { label: '미디어', value: 543 },
              ]}
              chartColors={[
                theme.palette.primary.main,
                theme.palette.info.main,
                theme.palette.warning.main,
                theme.palette.error.main,
              ]}
            />
          </Grid>
          <Grid item xs={12} md={6} lg={4}>
            <AppOrderTimeline
              title="STO 진행 절차"
              list={[...Array(5)].map((_, index) => ({
                id: faker.datatype.uuid(),
                title: [
                  '광고 등록',
                  'STO 생성',
                  'STO 모금',
                  '광고 입찰',
                  '수익 분배',
                ][index],
                type: `order${index + 1}`,
              }))}
            />
          </Grid>

<Grid item xs={12} md={6} lg={8}>
  <AppNewsUpdate
    title="인기 광고"
    list={[{id: 1,
      title: "전광판",
      description: "옥외광고",
      image: `/assets/images/covers/cover_1.jpg`,
      postedAt: faker.date.recent(),}, {id: 1,
        title: "전광판",
        description: "옥외광고",
        image: `/assets/images/covers/cover_1.jpg`,
        postedAt: faker.date.recent(),}, {id: 1,
          title: "전광판",
          description: "옥외광고",
          image: `/assets/images/covers/cover_1.jpg`,
          postedAt: faker.date.recent(),}, {id: 1,
            title: "전광판",
            description: "옥외광고",
            image: `/assets/images/covers/cover_1.jpg`,
            postedAt: faker.date.recent(),}, {id: 1,
      title: "전광판",
      description: "옥외광고",
      image: `/assets/images/covers/cover_1.jpg`,
      postedAt: faker.date.recent(),}, 
    ]}
  />
</Grid>

        </Grid>
      </Container>
    </>
  );
}
