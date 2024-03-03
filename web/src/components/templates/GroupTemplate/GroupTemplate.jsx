import { Box, Container, Stack } from "@mui/material";
import GroupTabs from "../../organisms/GroupTabs/GroupTabs";

function GroupTemplate({ children }) {
  return (
    <Box py={8}>
      <Container>
        <Stack spacing={4}>
          <GroupTabs />

          {children}
        </Stack>
      </Container>
    </Box>
  );
}

export default GroupTemplate;
