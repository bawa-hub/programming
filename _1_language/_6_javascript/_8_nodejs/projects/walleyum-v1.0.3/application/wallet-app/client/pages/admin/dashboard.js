import dynamic from 'next/dynamic';
import Head from 'next/head';
import React from 'react';
import { Table } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import Loader from '../../components/Loader';
import SidebarAdmin from '../../components/SidebarAdmin';
import UserHeaderAdmin from '../../components/UserHeaderAdmin';
import useDashboard from '../../data/useDashboard';
import { useUsersDashboard } from '../../data/useUsers';
import withAuthAdmin from '../../hoc/withAuthAdmin';

const Chart = dynamic(() => import('react-apexcharts'), { ssr: false });

const Dashboard = ({ userData, settings }) => {
  const { data, loading } = useDashboard();
  const { data: userList, loading: userLoading } = useUsersDashboard(1, 10);

  const { t } = useTranslation();

  if (loading || userLoading) {
    return <Loader height="100vh" />;
  }
  const options = {
    chart: {
      height: 350,
      type: 'area',
      foreColor: '#8C87C2',
      background: 'transparent',
      stacked: true,
      dropShadow: {
        enabled: true,
        enabledSeries: [0],
        top: -2,
        left: 2,
        blur: 5,
        opacity: 0.06
      },
      toolbar: {
        show: false,
      }
    },
    dataLabels: {
      enabled: false,
    },
    grid: {
      show: false,
      padding: {
        left: -5,
        right: 5
      }
    },
    stroke: {
      curve: 'smooth',
      width: 3
    },
    markers: {
      size: 0,
      strokeColor: '#fff',
      strokeWidth: 3,
      strokeOpacity: 1,
      fillOpacity: 1,
      hover: {
        size: 6
      }
    },
    xaxis: {
      type: 'category',
      categories: data?.labels,
    },
    legend: {
      position: 'top',
      horizontalAlign: 'left'
    },
    fill: {
      type: 'gradient',
      gradient: {
        shadeIntensity: 1,
        opacityFrom: 0.5,
        opacityTo: 0,
        stops: [0, 100, 100]
      }
    },
    theme: {
      mode: 'dark',
      monochrome: {
        enabled: false
      }
    }
  };
  const series = [{
    name: t('Deposits'),
    data: data?.last7DaysDeposits
  }, {
    name: t('Withdraws'),
    data: data?.last7DaysWithdraws
  }, {
    name: t('Exchanges'),
    data: data?.last7DaysExchanges
  }];
  return (
    <>
      <Head>
        <title>
          {t('Admin Panel')}
        </title>
      </Head>
      <div>
        <UserHeaderAdmin />
        <SidebarAdmin userData={userData} settings={settings} />
        <div className="content-body">
          <div className="container-fluid">
            <div className="row">
              <div className="col-xs-12 col">
                <div className="dashboard-tiles basic-card">
                  <h3>{data?.totalDeposits}</h3>
                  <p>{t('Total Deposits')}</p>
                </div>
              </div>
              <div className="col-xs-12 col">
                <div className="dashboard-tiles basic-card">
                  <h3>{data?.totalWithdrawn}</h3>
                  <p>{t('Total Withdraws')}</p>
                </div>
              </div>
              <div className="col-xs-12 col">
                <div className="dashboard-tiles basic-card">
                  <h3>{data?.totalExchanges}</h3>
                  <p>{t('Total Exchanges')}</p>
                </div>
              </div>
              <div className="col-xs-12 col">
                <div className="dashboard-tiles basic-card">
                  <h3>{data?.totalUsers}</h3>
                  <p>{t('Total Users')}</p>
                </div>
              </div>
              <div className="col-xs-12 col">
                <div className="dashboard-tiles basic-card">
                  <h3>{data?.totalMerchants}</h3>
                  <p>{t('Total Merchants')}</p>
                </div>
              </div>
            </div>
            <div className="row">
              <div className="col">
                <div className="basic-card mb-20">
                  <h4 className="box-title">
                    {t('Last 7 Days')}
                  </h4>
                  <Chart options={options} series={series} type="area" height={350} />
                </div>
              </div>
            </div>
            <div className="row">
              <div className="col">
                <div className="basic-card">
                  <h4 className="box-title">
                    {t('Recently Registered')}
                  </h4>
                  <Table striped hover className="dark-color" responsive>
                    <thead>
                      <tr>
                        <th scope="col">{t('ID')}</th>
                        <th scope="col">{t('Name')}</th>
                        <th scope="col">{t('Email')}</th>
                        <th scope="col">{t('Status')}</th>
                      </tr>
                    </thead>
                    <tbody>
                      {userList?.data?.map((history) => (
                        <tr key={history?.id}>
                          <td>
                            <strong>
                              #
                              {history?.id}
                            </strong>
                          </td>
                          <td>
                            <strong>
                              {history?.name}
                            </strong>
                          </td>
                          <td>
                            <strong>
                              {history?.email}
                            </strong>
                          </td>
                          <td>
                            <strong>
                              {history?.active ? t('Active') : t('Inactive')}
                            </strong>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </Table>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withAuthAdmin(Dashboard);
