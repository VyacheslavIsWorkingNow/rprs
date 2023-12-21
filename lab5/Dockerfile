FROM sequenceiq/hadoop-docker:latest

COPY input.txt /input.txt
COPY run_hadoop_job.sh /run_hadoop_job.sh

RUN chmod +x /run_hadoop_job.sh

CMD service ssh start && $HADOOP_HOME/sbin/start-dfs.sh && $HADOOP_HOME/sbin/start-yarn.sh && /run_hadoop_job.sh
