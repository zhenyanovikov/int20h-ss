import { Box, Container, Stack } from "@mui/material";
import AdminTabs from "../../organisms/AdminTabs/AdminTabs";

function AdminTemplate({ children }) {
  return (
    <Box py={8}>
      <Container>
        <Stack spacing={4}>
          <AdminTabs />

          {children}
        </Stack>
      </Container>
    </Box>
  );
}

export default AdminTemplate;
