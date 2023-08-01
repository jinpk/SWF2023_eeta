import {useState, useEffect} from 'react';
import { Helmet } from 'react-helmet-async';
import { Container, Typography, Stack, Button } from '@mui/material';
import { SigningStargateClient } from '@cosmjs/stargate';
import BillboardList from '../sections/@dashboard/billboard/BillboardList';
import Iconify from '../components/iconify';

export default function BillboardPage() {
  const [billboards, setBillboards] = useState([]);

  useEffect(() => {
    async function fetchBillboards() {
      try {
        const response = await fetch('http://3.37.36.76:1317/eeta/billboards');
        const data = await response.json();
        setBillboards(data.billboard);
      } catch (error) {
        console.error('Error fetching billboards:', error);
      }
    }

    fetchBillboards();
  }, []);

  const handleSign = async () => {
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
    const signed = response.signed_doc;
            const {signature} = response;
            const client = new SigningStargateClient('http://3.37.36.76:1317');
            console.log(response);
            client.broadcastTx({
              msg: signed.msgs,
              fee: signed.fee,
              signatures: [signature],
              memo: signed.memo,
            });
  };

  return (
    <>
      <Helmet>
        <title> Dashboard: Blog | Minimal UI </title>
      </Helmet>

      <Container>
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={5}>
          <Typography variant="h4" gutterBottom>
            광고
          </Typography>
          <Button variant="contained" startIcon={<Iconify icon="eva:plus-fill" />} onClick={handleSign}>
            광고 추가
          </Button>
        </Stack>

        
        <BillboardList billboards={billboards} />
      </Container>
    </>
  );
}
