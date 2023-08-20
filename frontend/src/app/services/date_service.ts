export const monthsList = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December',
];

export const formatDate = (datetime: number): string => {
  const dateObj: Date = new Date(datetime);
  const date: number = dateObj.getDate();
  const month: string = monthsList[dateObj.getMonth()];
  const year: number = dateObj.getFullYear();
  let hour: number = dateObj.getHours();
  const minutes: number = dateObj.getMinutes();
  let isAm = true;
  if (hour >= 12) {
    isAm = false;
    hour -= 12;
  }
  const formattedDate = `${hour}:${minutes} ${
    isAm ? 'AM' : 'PM'
  }, ${date} ${month} ${year}`;
  return formattedDate;
};
