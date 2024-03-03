import { useQueryClient, useQuery, useMutation } from "@tanstack/react-query";
import { enqueueSnackbar } from "notistack";
import { useTranslation } from "react-i18next";
import apiClient from "./apiClient";
import {
  STUDENTS_URL,
  INVITE_STUDENT_URL,
  STUDENTS_KEY,
} from "../constants/student";

async function getStudents(params) {
  const res = await apiClient.get(STUDENTS_URL, { params });
  return res.data;
}

async function inviteStudent(student) {
  const res = await apiClient.post(INVITE_STUDENT_URL, student);
  return res.data;
}

export function useGetStudents(params) {
  return useQuery({
    queryKey: [STUDENTS_KEY, params],
    queryFn: () => getStudents(params),
  });
}

export function useInviteStudent() {
  const { t } = useTranslation();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: inviteStudent,
    onSuccess: () => {
      enqueueSnackbar(t("snackbars.inviteStudent.success"), {
        variant: "success",
      });
      queryClient.invalidateQueries({ queryKey: [STUDENTS_KEY] });
    },
    onError: () => {
      enqueueSnackbar(t("snackbars.inviteStudent.error"), { variant: "error" });
    },
  });
}
