using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;


public delegate void CallbackTouch(int x, int y);
    
public class Handler
{
    #region private members
    private Thread tcpListenerThread;
    #endregion

    private static Server instance = null;
    CallbackTouch callbackTouch;

    void StartReading()
    {
        Debug.Log("Start Server");
        instance = this;

        // Start TcpServer background thread
        tcpListenerThread = new Thread(new ThreadStart(ListenForIncommingRequest));
        tcpListenerThread.IsBackground = true;
        tcpListenerThread.Start();
    }

    // Update is called once per frame
    void Update()
    {

    }

    public static Server Instance
    {
        get
        {
            if(instance == null)
            {
                return null;
            }
            return instance;
        }
    }

    public void SetTouchCallback(CallbackTouch callback)
    {
        if(callbackTouch == null)
        {
            callbackTouch = callback;
        } else
        {
            callbackTouch += callback;
        }
    }

    // Runs in background TcpServerThread; Handles incomming TcpClient requests
    private void ListenForIncommingRequest()
    {
        try
        {
            while (true)
            {
                // Get a stream object for reading
                using (NetworkStream stream = connectedTcpClient.GetStream())
                {
                    // Read incomming stream into byte array.
                    do
                    {
                        Byte[] bytesTypeOfService = new Byte[4];
                        Byte[] bytesPayloadLength = new Byte[4];

                        int lengthTypeOfService = stream.Read(bytesTypeOfService, 0, 4);
                        int lengthPayloadLength = stream.Read(bytesPayloadLength, 0, 4);

                        if (lengthTypeOfService <= 0 && lengthPayloadLength <= 0)
                        {
                            break;
                        }

                        // Reverse byte order, in case of big endian architecture
                        if (!BitConverter.IsLittleEndian)
                        {
                            Array.Reverse(bytesTypeOfService);
                            Array.Reverse(bytesPayloadLength);
                        }

                        int typeOfService = BitConverter.ToInt32(bytesTypeOfService, 0);
                        int payloadLength = BitConverter.ToInt32(bytesPayloadLength, 0);

                        Byte[] bytes = new Byte[payloadLength];
                        int length = stream.Read(bytes, 0, payloadLength);

                        HandleIncommingRequest(typeOfService, payloadLength, bytes);
                    } while (true);
                }
            }
        }
        catch (SocketException socketException)
        {
            Debug.Log("SocketException " + socketException.ToString());
        }
    }

    // Handle incomming request
    private void HandleIncommingRequest(int typeOfService, int payloadLength, byte[] bytes)
    {
        Debug.Log("=========================================");
        Debug.Log("Type of Service : " + typeOfService);
        Debug.Log("Payload Length  : " + payloadLength);
        switch (typeOfService)
        {
            case 0:
                TouchHandler(payloadLength, bytes);
                break;
        }
    }

    // Handle Touch Signal
    private void TouchHandler(int payloadLength, byte[] bytes)
    {
        Debug.Log("Execute Touch Handler");
        int x_axis = BitConverter.ToInt32(bytes, 0);
        int y_axis = BitConverter.ToInt32(bytes, 4);
        Debug.Log("X axis     : " + x_axis);
        Debug.Log("Y axis     : " + y_axis);
        if(callbackTouch != null)
        {
            callbackTouch(x_axis, y_axis);
        }
    }
}
