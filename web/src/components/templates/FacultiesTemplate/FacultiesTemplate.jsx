import { Box, Stack, CircularProgress } from "@mui/material";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";
import Empty from "../../molecules/Empty/Empty";

function FacultiesTemplate({ facultiesData, isGetFacultiesLoading }) {
  return (
    <AdminTemplate>
      {isGetFacultiesLoading || !facultiesData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <>{!facultiesData.length ? <Empty /> : <Box />}</>
      )}
  </AdminTemplate>
  );
}

export default FacultiesTemplate;
