FROM ollama/ollama:0.5.4
RUN /bin/sh -c "/bin/ollama serve & sleep 1 && ollama pull qwen2.5:0.5b"
ENTRYPOINT ["/bin/ollama"]
EXPOSE 11434
CMD ["serve"]
