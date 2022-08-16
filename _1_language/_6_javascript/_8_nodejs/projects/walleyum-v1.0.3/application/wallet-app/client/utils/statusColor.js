const statusColor = (status) => {
  if (status === 'pending') {
    return 'info';
  } if (status === 'success') {
    return 'success';
  } if (status === 'failed') {
    return 'danger';
  }
  return 'success';
};
export default statusColor;
