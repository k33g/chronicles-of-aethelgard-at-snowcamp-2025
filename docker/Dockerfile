FROM ollama/ollama:0.5.4
RUN /bin/sh -c "/bin/ollama serve & sleep 1 && ollama pull qwen2.5:0.5b"
RUN /bin/sh -c "/bin/ollama serve & sleep 1 && ollama pull qwen2.5:1.5b"
RUN /bin/sh -c "/bin/ollama serve & sleep 1 && ollama pull snowflake-arctic-embed:33m"
ENTRYPOINT ["/bin/ollama"]
EXPOSE 11434
CMD ["serve"]
